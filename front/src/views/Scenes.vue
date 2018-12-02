<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th>Name</th>
          <th>EntranceMessage</th>
          <th>VROMStart</th>
          <th>VROMEnd</th>
          <th>Rooms</th>
          <th>Data size</th>
          <th>Actions</th>
        </tr>
      </thead>

      <tbody>
        <tr
          v-for="scene, _, index in scenes"
          :key="index"
          v-if="scene.VROMStart > 0"
        >
          <td>{{ scene.Name }}</td>
          <td>{{ scene.EntranceMessage }}</td>
          <td>{{ scene.VROMStart | hex(8) }}</td>
          <td>{{ scene.VROMEnd | hex(8) }}</td>
          <td>{{ scene.Rooms === null ? 0 : scene.Rooms.length }}</td>
          <td>{{ scene.VROMEnd - scene.VROMStart | humanizeBytes }}</td>
          <td>
            <div class="field is-grouped">
              <p class="control">
                <a class="button" :href="'/api/files/' + scene.VROMStart | apiURI">Download</a>
              </p>
              <p class="control">
                <RouterLink
                  class="button is-primary"
                  :to="{name: 'SceneDetail', params: {start: scene.VROMStart}}"
                >Details</RouterLink>
              </p>
            </div>
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
