<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th>Name</th>
          <th>Type</th>
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
          <td>{{ file.Type }}</td>
          <td>{{ file.VROMStart | hex(8) }}</td>
          <td>{{ file.VROMEnd | hex(8) }}</td>
          <td>{{ file.VROMEnd - file.VROMStart | humanizeBytes }}</td>
          <td>
            <div class="field is-grouped">
              <p class="control">
                <a class="button" :href="'/api/files/' + file.VROMStart | apiURI">Download</a>
              </p>
              <p v-if="file.Type == 'scene'" class="control">
                <RouterLink
                  class="button is-primary"
                  :to="{name: 'SceneDetail', params: {start: file.VROMStart}}"
                >Details</RouterLink>
              </p>
              <p v-if="file.Type == 'room'" class="control">
                <RouterLink
                  class="button is-primary"
                  :to="{name: 'RoomDetail', params: {start: file.VROMStart}}"
                >Details</RouterLink>
              </p>
            </div>
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
