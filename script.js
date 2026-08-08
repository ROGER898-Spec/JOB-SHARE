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