<template>
  <BreadCrumb />
  <header class="bg-blue-600 text-white p-10 text-center">
  <h1 class="text-4xl font-bold">Our Services</h1>
</header>
  <div class="max-w-7xl mx-auto p-6">
    <h2 class="text-xl text-center mb-10">
      Reliable Solutions for Data Management
    </h2>

    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
      <!-- Data Backup Service Card -->
      <div class="bg-white shadow-lg rounded-lg p-6">
        <h3 class="text-2xl font-semibold mb-2">Data Backup 🗄️</h3>
        <p class="mb-4">
          Ensure your data is safe with our automated backup solutions.
        </p>
        <ul class="list-disc list-inside mb-4">
          <li>Scheduled backups</li>
          <li>Incremental backup options</li>
          <li>Easy data restoration</li>
        </ul>
      </div>

      <!-- Data Restore Service Card -->
      <div class="bg-white shadow-lg rounded-lg p-6">
        <h3 class="text-2xl font-semibold mb-2">Data Restore 🔄</h3>
        <p class="mb-4">
          Quick and reliable data restoration services to retrieve lost or
          corrupted data.
        </p>
        <ul class="list-disc list-inside mb-4">
          <li>Fast recovery times</li>
          <li>Support for various formats</li>
          <li>Comprehensive restoration plans</li>
        </ul>
      </div>

      <!-- Data Migration Service Card -->
      <div class="bg-white shadow-lg rounded-lg p-6">
        <h3 class="text-2xl font-semibold mb-2">Data Migration 🚀</h3>
        <p class="mb-4">
          Seamlessly migrate your data between platforms with our expert
          migration services.
        </p>
        <ul class="list-disc list-inside mb-4">
          <li>Platform-to-platform migration</li>
          <li>Minimal disruption during migration</li>
          <li>Data integrity checks</li>
        </ul>
      </div>
    </div>

    <div class="mt-10">
      <h3 class="text-2xl font-bold mb-4">Why Choose Us?</h3>
      <p class="mb-4">
        Expert Team, Cutting-Edge Technology, Customer-Centric Approach
      </p>
    </div>

    <div class="mt-10">
      <h3 class="text-2xl font-bold mb-4">Testimonials</h3>
      <blockquote class="border-l-4 border-gray-400 pl-4 italic mb-2">
        “The data backup service has saved us countless hours of work. Highly
        recommended!”
      </blockquote>
      <blockquote class="border-l-4 border-gray-400 pl-4 italic">
        “Thanks to their migration service, we switched platforms without losing
        any data!”
      </blockquote>
    </div>

    <div class="mt-10 text-center">
      <h3 class="text-2xl font-bold mb-4">Contact Us</h3>
      <p class="mb-4">
        Ready to secure your data?
        <a href="#" class="text-blue-500 hover:underline"
          >Get in touch with us today!</a
        >
      </p>
      <button
        class="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600"
      >
        Request a Quote
      </button>
    </div>
  </div>
</template>
