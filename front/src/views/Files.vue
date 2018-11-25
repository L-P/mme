<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th>Name</th>
          <th>VROMStart</th>
          <th>VROMEnd</th>
          <th>Data size</th>
          <th>Actions</th>
        </tr>
      </thead>

      <tbody>
        <tr
          v-for="file, _, index in files"
          :key="index"
        >
          <td>{{ file.Name }}</td>
          <td>{{ file.VROMStart | hex(8) }}</td>
          <td>{{ file.VROMEnd | hex(8) }}</td>
          <td>{{ file.VROMEnd - file.VROMStart | humanizeBytes }}</td>
          <td>
            <a :href="'/api/files/' + file.VROMStart | apiURI">
              DL
            </a>
          </td>
        </tr>
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
