<template>
  <header class="flex items-center justify-between h-12 text-xl bg-blue-600">
    <div
      class="bg-blue-600 ml-1 text-white flex flex-col items-center py-1 text-sm h-full"
    >
      <p>AP</p>
      <p>Technologies</p>
    </div>
    <div class="flex items-center">
      <router-link
        to="/home"
        class="list-none active:text-[3c48b00] px-3 text-white cursor-pointer"
        >Home</router-link
      >
      <router-link
        to="aboutus"
        class="list-none active:text-[3c48b00] px-3 text-white cursor-pointer"
      >
        About Us
      </router-link>
      <router-link
        to="contact"
        class="list-none active:text-[3c48b00] px-3 text-white cursor-pointer"
      >
        Contact
      </router-link>

      <router-link
        to="certification"
        class="list-none active:text-[3c48b00] px-3 text-white cursor-pointer"
      >
        Certification
      </router-link>
      <!-- <router-link
        to="services"
        class="list-none active:text-[3c48b00] px-3 text-white cursor-pointer"
      >
        Services
      </router-link> -->
      <router-link to="services">
        <el-dropdown split-button type="primary">
          Services
          <template #dropdown>
            <el-dropdown-menu>
              <router-link to="services/data-backup">
                <el-dropdown-item>Data Backup</el-dropdown-item>
              </router-link>
              <router-link to="services/data-restore">
                <el-dropdown-item>Data Restore</el-dropdown-item>
              </router-link>
              <router-link to="services/data-migration">
                <el-dropdown-item>Data Migration</el-dropdown-item>
              </router-link>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </router-link>
    </div>
    <button class="border mr-1 text-base rounded border-white text-white px-1">
      Get In Touch
    </button>
  </header>
  <hr />
</template>

<style>
.el-button--primary {
  --el-button-bg-color: #2563eb !important;
  --el-button-outline-color: #2563eb !important;
  --el-button-active-color: #2563eb !important;
  border: none !important;
  padding: 0 6px !important;
  font-size: 20px !important;
  font-weight: 400 !important;
  margin-top: -2px !important;
}
.el-button:hover {
  background-color: #2563eb !important;
}
</style>
