import Swiper, { Pagination, Navigation } from 'swiper';
import 'swiper/swiper-bundle.css';
import { Autoplay } from 'swiper';

const customPaginationSourceCode = `
  <div class="swiper swiper-with-custom-pagination">
    <div class="swiper-wrapper">
      <!-- Swiper slides go here -->
    </div>
    <div class="swiper-pagination swiper-pagination-custom"></div>
  </div>
`;

const customSwiperSourceCode = `
  <div class="swiper swiper-custom">
    <div class="swiper-wrapper">
      <!-- Swiper slides go here -->
    </div>
    <div class="swiper-button-next"></div>
    <div class="swiper-button-prev"></div>
    <div class="swiper-pagination"></div>
  </div>
`;

// Create CodeViewer
const customPaginationCodeViewer = createCodeViewer('#custom-pagination-code-viewer', customPaginationSourceCode);
const customSwiperCodeViewer = createCodeViewer('#custom-caroursel-code-viewer', customSwiperSourceCode);

// Render CodeViewer
customPaginationCodeViewer.render();
customSwiperCodeViewer.render();

// Initialize Swiper instances
document.addEventListener("DOMContentLoaded", function() {
  const customPaginationSwiper = new Swiper('.swiper-with-custom-pagination', {
    modules: [Pagination],
    pagination: {
      el: '.swiper-pagination',
      clickable: true,
      renderBullet: function (index, className) {
        return '<span class="' + className + '">' + (index + 1) + '</span>';
      },
    },
  });

  const customSwiper = new Swiper('.swiper-custom', {
    loop: true,
    modules: [Pagination, Navigation],
    pagination: {
      el: '.swiper-pagination',
      type: 'bullets',
      clickable: true,
    },
    navigation: {
      nextEl: '.swiper-button-next',
      prevEl: '.swiper-button-prev',
    },
  });
});
