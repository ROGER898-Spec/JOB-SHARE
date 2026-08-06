// Waitlist form — demo only, belum terhubung ke backend
  const form = document.getElementById('waitlistForm');
  const msg = document.getElementById('waitlistMsg');
  form.addEventListener('submit', function(e){
    e.preventDefault();
    const email = document.getElementById('email').value;
    if(email){
      msg.textContent = 'Terima kasih! ' + email + ' sudah masuk daftar tunggu.';
      form.reset();
    }
  });

  // Mobile menu toggle (placeholder — link list still hidden on small screens for now)
  const menuToggle = document.querySelector('.menu-toggle');
  const navLinks = document.querySelector('.nav-links');
  menuToggle.addEventListener('click', function(){
    const isOpen = navLinks.style.display === 'flex';
    navLinks.style.display = isOpen ? 'none' : 'flex';
    navLinks.style.flexDirection = 'column';
    navLinks.style.position = 'absolute';
    navLinks.style.top = '64px';
    navLinks.style.left = '0';
    navLinks.style.right = '0';
    navLinks.style.background = 'var(--paper-raised)';
    navLinks.style.padding = '20px 28px';
    navLinks.style.borderBottom = '1px solid var(--line)';
    navLinks.style.gap = '16px';
  });