package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "net/smtp"
  "os"
  "time"
)

type Contact struct {
  Name    string `json:"name"`
  Email   string `json:"email"`
  Message string `json:"message"`
}

func sendSMTP(to, host, port, user, pass string, c Contact) error {
  auth := smtp.PlainAuth("", user, pass, host)
  body := fmt.Sprintf("From: %s <%s>\r\nTo: %s\r\nSubject: Contact form: %s\r\n\r\n%s",
    c.Name, c.Email, to, c.Name, c.Message)
  addr := host + ":" + port
  return smtp.SendMail(addr, auth, user, []string{to}, []byte(body))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "method", http.StatusMethodNotAllowed)
    return
  }
  var c Contact
  dec := json.NewDecoder(r.Body)
  if err := dec.Decode(&c); err != nil {
    http.Error(w, "bad json", http.StatusBadRequest)
    return
  }

  // Basic validation
  if c.Email == "" || c.Message == "" {
    http.Error(w, "missing fields", http.StatusBadRequest)
    return
  }

  // Try SMTP if env provided
  smtpHost := os.Getenv("SMTP_HOST")
  smtpPort := os.Getenv("SMTP_PORT")
  smtpUser := os.Getenv("SMTP_USER")
  smtpPass := os.Getenv("SMTP_PASS")
  toEmail := os.Getenv("TO_EMAIL")

  if smtpHost != "" && toEmail != "" {
    if smtpPort == "" { smtpPort = "587" }
    err := sendSMTP(toEmail, smtpHost, smtpPort, smtpUser, smtpPass, c)
    if err != nil {
      log.Printf("smtp send error: %v", err)
      http.Error(w, "send failed", http.StatusInternalServerError)
      return
    }
    json.NewEncoder(w).Encode(map[string]string{"message":"sent"})
    return
  }

  // Fallback: append to local file in current dir
  f, err := os.OpenFile("contact_messages.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil {
    log.Printf("open file error: %v", err)
    http.Error(w, "server error", http.StatusInternalServerError)
    return
  }
  defer f.Close()
  entry := fmt.Sprintf("%s\nName: %s\nEmail: %s\nMessage:\n%s\n---\n",
    time.Now().Format(time.RFC3339), c.Name, c.Email, c.Message)
  if _, err := f.WriteString(entry); err != nil {
    log.Printf("write error: %v", err)
    http.Error(w, "server error", http.StatusInternalServerError)
    return
  }
  json.NewEncoder(w).Encode(map[string]string{"message":"logged"})
}

func main() {
  http.HandleFunc("/contact", contactHandler)
  port := os.Getenv("PORT")
  if port == "" { port = "8080" }
  addr := ":" + port
  log.Printf("contact service listening %s", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    log.Fatalf("server: %v", err)
  }
}
