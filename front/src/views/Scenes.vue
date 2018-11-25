<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th>Name</th>
          <th>EntranceMessage</th>
          <th>VROMStart</th>
          <th>VROMEnd</th>
          <th>Data size</th>
          <th>Actions</th>
        </tr>
      </thead>

      <tbody>
        <tr
          v-for="scene, _, index in scenes"
          :key="index"
        >
          <td>{{ scene.Name }}</td>
          <td>{{ scene.EntranceMessage }}</td>
          <td>{{ scene.VROMStart | hex(8) }}</td>
          <td>{{ scene.VROMEnd | hex(8) }}</td>
          <td>{{ scene.VROMEnd - scene.VROMStart | humanizeBytes }}</td>
          <td>
            <a :href="'/api/files/' + scene.VROMStart | apiURI">
              DL
            </a>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script src="./Scenes.js"></script>

<style scoped>
table {
  width: 100%;
}
</style>
