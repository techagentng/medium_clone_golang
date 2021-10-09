
let modalBt = document.querySelector('.slt');
let modalBg = document.querySelector('.modal-bg');
modalBt.addEventListener('click', function(){
    modalBg.classList.add('bg-active');
})
let closeModal = document.querySelector('.modal-close');
closeModal.addEventListener('click', function(){
    modalBg.classList.remove('bg-active')
})