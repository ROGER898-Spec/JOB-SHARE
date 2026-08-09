// components.js
// Navbar dan footer JOBnesia — dipakai bersama di semua halaman.

const NAV_LINKS = [
  { label: 'Beranda', href: 'index.html' },
  { label: 'Cara Kerja', href: 'cara-kerja.html' },
  { label: 'Jasa', href: 'jasa.html' },
  { label: 'About Us', href: 'about.html' },
  { label: 'Kontak', href: 'kontak.html' }
];

function renderHeader(){
  const header = document.getElementById('site-header');
  if(!header) return;

  const linksHtml = NAV_LINKS.map(function(link){
    const isActive = window.location.pathname.endsWith(link.href) || (window.location.pathname.endsWith('/') && link.href === 'index.html');
    const style = isActive ? ' style="color:var(--teal); font-weight:700;"' : '';
    return '<li><a href="' + link.href + '"' + style + '>' + link.label + '</a></li>';
  }).join('');

  header.innerHTML =
    '<nav class="nav wrap" style="padding-left:0; padding-right:0;">' +
      '<a href="index.html" class="logo" style="text-decoration:none;"><span class="dot"></span>JOBnesia</a>' +
      '<ul class="nav-links" id="nav-links">' + linksHtml + '</ul>' +
      '<div class="nav-cta">' +
        '<a href="index.html#cta" class="btn btn-ghost btn-sm mobile-hide">Masuk</a>' +
        '<a href="index.html#cta" class="btn btn-primary btn-sm"><span class="long">Gabung&nbsp;</span>Gratis</a>' +
      '</div>' +
      '<button class="menu-toggle" id="mobile-menu" aria-label="Buka menu"><span></span><span></span><span></span></button>' +
    '</nav>';

  // Mobile menu toggle — menggunakan CSS class agar lebih bersih dan mendukung animasi
  const menuToggle = header.querySelector('#mobile-menu');
  const navLinksEl = header.querySelector('#nav-links');
  const navItems = navLinksEl.querySelectorAll('a');

  menuToggle.addEventListener('click', function(){
    navLinksEl.classList.toggle('active');
    menuToggle.classList.toggle('active');
  });

  // Memastikan menu tertutup saat hyperlink diklik di tampilan mobile
  navItems.forEach(item => {
    item.addEventListener('click', function() {
      navLinksEl.classList.remove('active');
      menuToggle.classList.remove('active');
    });
  });
}

function renderFooter(){
  const footer = document.getElementById('site-footer');
  if(!footer) return;

  footer.innerHTML =
    '<div class="wrap footer-inner">' +
      '<a href="index.html" class="logo" style="text-decoration:none;"><span class="dot"></span>JOBnesia</a>' +
      '<p>Dibuat untuk GEMASTIK XIX 2026 — Kompetisi Pengembangan Perangkat Lunak</p>' +
    '</div>';
}

document.addEventListener('DOMContentLoaded', function(){
  renderHeader();
  renderFooter();
});