// Hide header on scroll down, show on scroll up. Small, dependency-free.
(function(){
  let lastY = 0;
  const header = document.getElementById('site-header');
  if(!header) return;
  const threshold = 10;
  window.addEventListener('scroll', ()=>{
    const y = window.scrollY || window.pageYOffset;
    if (Math.abs(y - lastY) < threshold) return;
    if (y > lastY && y > 60) {
      header.style.transform = 'translateY(-100%)';
    } else {
      header.style.transform = '';
    }
    lastY = y;
  }, {passive:true});
})();
