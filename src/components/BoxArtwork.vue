<template>
    
    <div>
        <div class="capa-obra">
            <div class="column1">
                <h3>Lorem Ipsun</h3>
                <p>Autor Desconhecido</p>
            </div>
            <button class="detalhes"><svg style="fill:#fff" width="1.5rem" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M256 80c0-17.7-14.3-32-32-32s-32 14.3-32 32V224H48c-17.7 0-32 14.3-32 32s14.3 32 32 32H192V432c0 17.7 14.3 32 32 32s32-14.3 32-32V288H400c17.7 0 32-14.3 32-32s-14.3-32-32-32H256V80z"/></svg></button>
        </div>
    </div>
    
</template>

<script>

export default {
  name: 'BoxArtwork',
  data() {
    return {
    };
  },
};

</script>

<style scoped>
.capa-obra{
    font-family: 'Poppins';
    background-image: url("../assets/obras/obra-1.jpg");
    background-size: cover;
    background-color: rgb(0, 0, 0, 0.5);
    display: flex;
    flex-direction: row;
    width: 250px;
    height: 250px;
    align-items: center;
    padding-top: 180px;
    justify-content: space-around;
    border-radius: 20px;
    position: relative;
}
.detalhes{
    width: min-content;
    height: min-content;
}
h3{
    color: white;
    font-size: 1.3rem;
    line-height: 1.2rem;
}
p{
    color: white;
    font-size: 0.8rem;
    line-height: 1rem;
}
.detalhes{
    background-color: var(--vt-c-orange);
    border: 0px;
    border-radius: 100%;
    color: white;
    font-size: 1rem;
    padding: 0px;
    width: 40px;
    height: 40px;
}
.detalhes:hover{
    background-color: var(--vt-c-indigo);
    transition: 0.3s;
}
</style>



