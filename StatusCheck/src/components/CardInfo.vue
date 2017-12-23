<template>
  <div class="tile is-child">
  <div class="card">
  <header class="card-header">
  <p class="card-header-title">
  {{url}}
  </p>
  </header>
  <div class="card-content">
  <div class="content">
  Uptime (Last Month): {{last}}
  </div>
  <div class="content">
  Live: {{good}}
  <div class="">
  <button class="button is-fullwidth">Visit</button>
  </div>
  </div>
  </div>

  </div>
  </div>
</template>

<script>
export default {
  name: 'CardInfo',
  props: ['url', 'last', 'good'],
  data () {
    return {
    }

  }
}
</script>
