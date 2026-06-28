// Sends contact form via fetch to /contact
document.addEventListener('DOMContentLoaded', function(){
  const form = document.getElementById('contact-form');
  if(!form) return;
  form.addEventListener('submit', async function(e){
    e.preventDefault();
    const data = new FormData(form);
    const body = {};
    data.forEach((v,k)=> body[k]=v);
    try {
      const res = await fetch('/contact', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body)
      });
      const json = await res.json();
      alert(json.message || 'Sent');
      form.reset();
    } catch(err){
      alert('Send failed');
    }
  });
});
