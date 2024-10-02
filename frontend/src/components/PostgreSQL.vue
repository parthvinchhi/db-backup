<template>
    <div class="max-w-lg mx-auto p-6 bg-white rounded-lg shadow-md">
      <h2 class="text-2xl font-bold mb-4">Restore Database</h2>
      <form @submit.prevent="handleSubmit">
        <div class="mb-4">
          <label for="host" class="block mb-1 font-semibold">Host:</label>
          <input
            type="text"
            id="host"
            v-model="form.host"
            class="border border-gray-300 rounded p-2 w-full"
            required
          />
        </div>
  
        <div class="mb-4">
          <label for="user" class="block mb-1 font-semibold">User:</label>
          <input
            type="text"
            id="user"
            v-model="form.user"
            class="border border-gray-300 rounded p-2 w-full"
            required
          />
        </div>
  
        <div class="mb-4">
          <label for="port" class="block mb-1 font-semibold">Port:</label>
          <input
            type="number"
            id="port"
            v-model="form.port"
            class="border border-gray-300 rounded p-2 w-full"
            required
          />
        </div>
  
        <div class="mb-4">
          <label for="password" class="block mb-1 font-semibold">Password:</label>
          <input
            type="password"
            id="password"
            v-model="form.password"
            class="border border-gray-300 rounded p-2 w-full"
            required
          />
        </div>
  
        <div class="mb-4">
          <label for="name" class="block mb-1 font-semibold">Database Name:</label>
          <input
            type="text"
            id="name"
            v-model="form.name"
            class="border border-gray-300 rounded p-2 w-full"
            required
          />
        </div>
  
        <div class="mb-4">
          <label for="sslmode" class="block mb-1 font-semibold">SSL Mode:</label>
          <select
            id="sslmode"
            v-model="form.sslmode"
            class="border border-gray-300 rounded p-2 w-full"
            required
          >
            <option value="disable">Disable</option>
            <option value="require">Require</option>
            <option value="verify-ca">Verify CA</option>
            <option value="verify-full">Verify Full</option>
          </select>
        </div>
  
        <button
          type="submit"
          class="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600"
        >
          Restore
        </button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        form: {
          host: '',
          user: '',
          port: '',
          password: '',
          name: '',
          sslmode: 'disable',
        },
      };
    },
    methods: {
      handleSubmit() {
        // Handle form submission logic here
        console.log('Form submitted:', this.form);
        // You might want to send this data to your API here
        alert('Database restored with the following details:\n' + JSON.stringify(this.form));
      },
    },
  };
  </script>
  
  <style scoped>
  /* Add any additional styles here if needed */
  </style>
  