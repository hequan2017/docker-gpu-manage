const navLinks = document.querySelectorAll('.nav a');
const sections = [...document.querySelectorAll('section[id]')];
const scrollBtn = document.querySelector('[data-scroll]');

const setActiveLink = () => {
  const scrollPos = window.scrollY + 200;
  sections.forEach((section) => {
    const id = section.getAttribute('id');
    const link = document.querySelector(`.nav a[href="#${id}"]`);
    if (!link) return;
    if (scrollPos >= section.offsetTop && scrollPos < section.offsetTop + section.offsetHeight) {
      navLinks.forEach((item) => item.classList.remove('active'));
      link.classList.add('active');
    }
  });
};

navLinks.forEach((link) => {
  link.addEventListener('click', (event) => {
    event.preventDefault();
    const target = document.querySelector(link.getAttribute('href'));
    if (target) {
      target.scrollIntoView({ behavior: 'smooth' });
    }
  });
});

if (scrollBtn) {
  scrollBtn.addEventListener('click', () => {
    const target = document.querySelector(scrollBtn.dataset.scroll);
    if (target) {
      target.scrollIntoView({ behavior: 'smooth' });
    }
  });
}

window.addEventListener('scroll', setActiveLink);
window.addEventListener('load', setActiveLink);
