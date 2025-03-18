<template>
  <!-- Footer -->
  <footer class="text-center text-sm-start bg-body-tertiary text-muted">
    <!-- Section: Social media -->
    <section
      class="d-flex justify-content-center justify-content-lg-between"
    >
      <!-- Right -->
      <div>
        <a href="" class="me-4 text-reset">
          <i class="fab fa-facebook-f"></i>
        </a>
        <a href="" class="me-4 text-reset">
          <i class="fab fa-twitter"></i>
        </a>
        <a href="" class="me-4 text-reset">
          <i class="fab fa-google"></i>
        </a>
        <a href="" class="me-4 text-reset">
          <i class="fab fa-instagram"></i>
        </a>
        <a href="" class="me-4 text-reset">
          <i class="fab fa-linkedin"></i>
        </a>
        <a href="" class="me-4 text-reset">
          <i class="fab fa-github"></i>
        </a>
      </div>
      <!-- Right -->
    </section>
    <!-- Section: Social media -->

    <!-- Section: Links  -->
    <section class="">
      <div class="container px-5 text-center text-md-start mt-5">
        <!-- Grid row -->
        <div class="row mt-3">
          <!-- Grid column -->
          <div class="col-md-3 col-lg-4 col-xl-3 mx-auto mb-4">
            <!-- Content -->
            <img
              src="@/assets/logo-virtual-marrom.png"
              alt="Icons"
              class="img-seta"
            />
            <p class="my-4">Explore museus virtualmente em <br>todo Brasil</p>
          </div>
          <!-- Grid column -->

          <!-- Grid column -->
          <div class="col-md-2 col-lg-2 col-xl-2 mx-auto mb-4">
            <!-- Links -->
            <h6 class="text-uppercase fw-bold mb-4">Sobre</h6>
            <p>
              <a href="#!" class="text-reset">Museus</a>
            </p>
            <p>
              <a href="#!" class="text-reset">Obras</a>
            </p>
            <p>
              <a href="#!" class="text-reset">Exposições</a>
            </p>
          </div>
          <!-- Grid column -->

          <!-- Grid column -->
          <div class="col-md-3 col-lg-2 col-xl-2 mx-auto mb-4">
            <!-- Links -->
            <h6 class="text-uppercase fw-bold mb-4">Parceiros</h6>
            <p>
              <a href="#!" class="text-reset">Crie um museu</a>
            </p>
            <p>
              <a href="#!" class="text-reset">Divulgue sua exposição</a>
            </p>
            <p>
              <a href="#!" class="text-reset">Convide um amigo</a>
            </p>
          </div>
          <!-- Grid column -->

          <!-- Grid column -->
          <div class="col-md-4 col-lg-3 col-xl-3 mx-auto mb-md-0 mb-4">
            <!-- Links -->
            <h6 class="text-uppercase fw-bold mb-4">Suporte</h6>
            <p> Entre em contato com o suporte:
              att@virtualartexplore.com</p>
            <p> <a href=""> Acesse nosso repositório </a></p>
          </div>
          <!-- Grid column -->
        </div>
        <!-- Grid row -->
      </div>
    </section>
    <!-- Section: Links  -->

    <!-- Copyright -->
    <div class="text-center p-3" style="background-color: rgba(0, 0, 0, 0.1); color: var(--vt-c-brown);">
      © 2024 Copyright:
      <a class="text-reset fw-bold" href="https://github.com/gustavoataidez/virtualartexplorer"
        >Virtual Art Explorer.</a
      > Todos os direitos reservados.
    </div>
    <!-- Copyright -->
  </footer>
  <!-- Footer -->
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {
    scrollToTop() {
      window.scrollTo({
        top: 0,
        behavior: "smooth",
      });
    },
  },
};
</script>

<style scoped>
.container {
  font-family: "Poppins", sans-serif;
}
footer {
  background-color: var(--vt-c-header) !important;
  width: 100%;
  z-index: 8;
  position: relative;
}

.img-seta{
  width: 100%;
}

p, a, h6 {
  color: var(--vt-c-brown);
}
</style>
