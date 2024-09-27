<template>
  <header class="flex items-center justify-between h-12 text-xl bg-blue-600">
    <div class="bg-blue-600 ml-1 text-white flex flex-col items-center py-1 text-sm h-full">
      <p>AP</p>
      <p>Technologies</p>
    </div>
    <div>
      <router-link
        to="/home"
        class="list-none active:text-[3c48b00]  px-3 text-white cursor-pointer"
        >Home</router-link
      >
      <router-link
        to="aboutus"
        class="list-none active:text-[3c48b00]  px-3 text-white cursor-pointer"
      >
        About Us
      </router-link>
      <router-link
        to="ourproducts"
        class="list-none active:text-[3c48b00]  px-3 text-white cursor-pointer"
      >
        Contact
      </router-link>

      <router-link
        to="certification"
        class="list-none active:text-[3c48b00]  px-3 text-white cursor-pointer"
      >
        Certification
      </router-link>
   
      <router-link
        to="contactus"
        class="list-none active:text-[3c48b00]  px-3 text-white cursor-pointer"
      >
        Services
      </router-link>
    </div>
    <button class="border mr-1 text-sm rounded border-white text-white px-1">Get In Touch</button>
  </header>
</template>
