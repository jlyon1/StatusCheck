<template>
  <div class="section">
    <div class="tile is-ancestor">
    <div class="tile is-parent is-vertical">
    <CardInfo url="example.com" good="true"></CardInfo>
    <CardInfo url="example.com" good="true"></CardInfo>
    <CardInfo url="example.com" good="true"></CardInfo>
    <CardInfo url="example.com" good="true"></CardInfo>
    </div>
    </div>
  </div>
</template>

<script>
import CardInfo from './CardInfo.vue'

export default {
  name: 'InfoSection',
  components:{
     'CardInfo':CardInfo
    },
  data () {
    return {
      test: 'Test!'
    }
  }
}
</script>
