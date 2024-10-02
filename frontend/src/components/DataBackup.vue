<template>
    <BreadCrumb />
  <div class="flex justify-center h-screen">
    <el-dropdown class="mt-3">
      <el-button type="primary">
        Select Your Data Type<el-icon class="el-icon--right"
          ><arrow-down
        /></el-icon>
      </el-button>
      <template #dropdown>
        <el-dropdown-menu>
          <router-link to="PostgreSQL">
            <el-dropdown-item>PostgreSQL</el-dropdown-item>
          </router-link>
          <router-link to="MySQL">
            <el-dropdown-item>MySQL </el-dropdown-item>
          </router-link>
          <router-link to="MongoDB">
            <el-dropdown-item>MongoDB</el-dropdown-item>
          </router-link>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </div>
</template>
