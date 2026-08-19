// script.js — logika khusus per halaman (di luar navbar/footer)

// Waitlist form — cuma ada di halaman index.html, jadi kita cek dulu
// supaya tidak error di halaman lain yang tidak punya form ini.
const form = document.getElementById('waitlistForm');
if(form){
  const msg = document.getElementById('waitlistMsg');
  form.addEventListener('submit', function(e){
    e.preventDefault();
    const email = document.getElementById('email').value;
    if(email){
      msg.textContent = 'Terima kasih! ' + email + ' sudah masuk daftar tunggu.';
      form.reset();
    }
  });
}

// FAQ accordion — cuma ada di halaman cara-kerja.html
const faqButtons = document.querySelectorAll('.faq-question');
faqButtons.forEach(function(btn){
  btn.addEventListener('click', function(){
    const item = btn.closest('.faq-item');
    const answer = item.querySelector('.faq-answer');
    const isOpen = item.classList.contains('open');

    // tutup item lain supaya tetap ringkas (accordion single-open)
    document.querySelectorAll('.faq-item.open').forEach(function(openItem){
      if(openItem !== item){
        openItem.classList.remove('open');
        openItem.querySelector('.faq-question').setAttribute('aria-expanded', 'false');
        openItem.querySelector('.faq-answer').setAttribute('aria-hidden', 'true');
      }
    });

    item.classList.toggle('open', !isOpen);
    btn.setAttribute('aria-expanded', String(!isOpen));
    if(answer){ answer.setAttribute('aria-hidden', String(isOpen)); }
  });
});

// Scroll-reveal untuk elemen ".reveal" — cuma ada di halaman cara-kerja.html
const revealEls = document.querySelectorAll('.reveal');
if(revealEls.length){
  if('IntersectionObserver' in window){
    const revealObserver = new IntersectionObserver(function(entries){
      entries.forEach(function(entry){
        if(entry.isIntersecting){
          entry.target.classList.add('is-visible');
          revealObserver.unobserve(entry.target);
        }
      });
    }, { threshold: 0.15 });
    revealEls.forEach(function(el){ revealObserver.observe(el); });
  } else {
    revealEls.forEach(function(el){ el.classList.add('is-visible'); });
  }
}