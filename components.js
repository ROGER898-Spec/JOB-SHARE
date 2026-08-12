// components.js
// Navbar dan footer JOBnesia — dipakai bersama di semua halaman.
// Ubah nav-links atau footer di SINI SAJA, otomatis berlaku di semua halaman
// (index.html, about.html, jasa.html, dst) tanpa perlu edit satu-satu.

const NAV_LINKS = [
  { label: 'Beranda', href: 'index.html' },
  { label: 'Cara Kerja', href: 'cara-kerja.html' },
  { label: 'Jasa', href: 'jasa.html' },
  { label: 'Kontak', href: 'kontak.html' },
  { label: 'About Us', href: 'about.html' },
];

function renderHeader(){
  const header = document.getElementById('site-header');
  if(!header) return;

  const linksHtml = NAV_LINKS.map(function(link){
    const isActive = window.location.pathname.endsWith(link.href);
    const style = isActive ? ' style="color:var(--teal); font-weight:700;"' : '';
    return '<li><a href="' + link.href + '"' + style + '>' + link.label + '</a></li>';
  }).join('');

  header.innerHTML =
    '<nav class="nav wrap" style="padding-left:0; padding-right:0;">' +
      '<a href="index.html" class="logo" style="text-decoration:none;"><span class="dot"></span>JOBnesia</a>' +
      '<ul class="nav-links">' + linksHtml + '</ul>' +
      '<div class="nav-cta">' +
        '<a href="login.html" class="btn btn-ghost btn-sm">Masuk</a>' +
        '<a href="register.html" class="btn btn-primary btn-sm"><span class="long">Gabung&nbsp;</span>Gratis</a>' +
      '</div>' +
      '<button class="menu-toggle" aria-label="Buka menu"><span></span><span></span><span></span></button>' +
    '</nav>';

  // Mobile menu toggle — dipasang ulang tiap kali header di-render
  const menuToggle = header.querySelector('.menu-toggle');
  const navLinksEl = header.querySelector('.nav-links');
  menuToggle.addEventListener('click', function(){
    const isOpen = navLinksEl.style.display === 'flex';
    navLinksEl.style.display = isOpen ? 'none' : 'flex';
    navLinksEl.style.flexDirection = 'column';
    navLinksEl.style.position = 'absolute';
    navLinksEl.style.top = '64px';
    navLinksEl.style.left = '0';
    navLinksEl.style.right = '0';
    navLinksEl.style.background = 'var(--paper-raised)';
    navLinksEl.style.padding = '20px 28px';
    navLinksEl.style.borderBottom = '1px solid var(--line)';
    navLinksEl.style.gap = '16px';
  });
}

function renderFooter(){
  const footer = document.getElementById('site-footer');
  if(!footer) return;

  footer.innerHTML =
    '<div class="wrap footer-inner">' +
      '<a href="index.html" class="logo" style="text-decoration:none;"><span class="dot"></span>JOBnesia</a>' +
      '<p>Proyek ini diharapkan mampu membantu mempermudah para UMKM untuk mencari freelancer yang mereka butuhkan</p>' +
    '</div>';
}

document.addEventListener('DOMContentLoaded', function(){
  renderHeader();
  renderFooter();
});