<template>
    <Header></Header>
    <div class="museu-page">
        <div class="sec-resume">
            <div class="capa-museu"></div>
            <div class="resume-desc">
                <span id="category-museum" class="badge rounded-pill text-bg-warning mb-2">Categoria</span>
                <h1>Museu Théo Brandão</h1>
                <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus totam sit nostrum dicta assumenda, officia atque facilis esse quasi eum perferendis vel aliquam reiciendis rem tempora quod doloribus maxime voluptas!</p>
                <h6 style="font-weight: 600; font-size: 1rem;">Mais informações</h6>
                <p style="line-height: 1.5rem;">Lorem ipsum dolor sit amet consectetur adipisicing elit. Illo aperiam eos eligendi explicabo </p>
            <div id="endereco" class="d-flex flex-row align-items-center">
                <img style="width: 60px; margin-right: 10px ;" src="../assets/location.png" alt="">
                <p style="line-height: 1.5rem; margin-bottom: 0rem;font-weight: 500;">Av. da Paz, 1490, Centro<br><span id="city-museum">Maceió</span> - <span id="state-museum">AL</span></p>
            </div>
            </div>
            
            </div>
        <div class="sec-obras">
            <span style="font-weight: 600;color: var(--vt-c-brown);line-height: 1rem;">Explore mais</span>
            <h2 style="line-height: 1.5rem;">Obras do Museu</h2>
            <div class="grade-obras">
            <BoxArtwork></BoxArtwork>
            <BoxArtwork></BoxArtwork>
            <BoxArtwork></BoxArtwork>
            <BoxArtwork></BoxArtwork>
            </div>
        </div>
    </div>
    <Footer></Footer>
</template>

<script>
import Header from "./HeaderPage.vue";
import Footer from "./Footer.vue";
import BoxArtwork from "./BoxArtwork.vue";


export default {
  components: {BoxArtwork,Header,Footer},
  data() {
    return {
    };
  },
};

</script>


<style>
.museu-page{
    font-family: 'Poppins';
    width: 90vw;
    margin-left: 50%;
    transform: translate(-50%,0);
    padding: 50px;
    max-width: 1200px;
    color: var(--vt-c-brown);
    display: flex;
    flex-direction: column;
    gap: 50px;
}
.sec-resume{
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    gap: 50px;
}
.resume-desc{
    width: 50%;
}
.resume-desc p{
    font-size: 0.9rem;
}
.capa-museu{
    width: 50%;
    height: 400px;
    border-radius: 50px;
    background-image: url('../assets/museus/museu-theo-brandao.jpg');
    background-repeat: no-repeat;
    background-size: cover;
}
.grade-obras{
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: center;
    gap: 30px;
    margin: 40px 0;
}
</style>