<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th>Name</th>
          <th>VROMStart</th>
          <th>VROMEnd</th>
        </tr>
      </thead>

      <tbody>
          <tr v-for="file in files">
            <td>{{ file.Name }}</td>
            <td>{{ file.VROMStart | hex(8) }}</td>
            <td>{{ file.VROMEnd | hex(8) }}</td>
          </tr>
        </v-for>
      </tbody>

    </table>
  </div>
</template>

<script src="./Files.js"></script>

<style scoped>
table {
  width: 100%;
}
</style>
