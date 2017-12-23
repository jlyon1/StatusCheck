<template>
  <div class="section">
    <div class="tile is-ancestor">
    <div class="tile is-parent is-vertical">
    <CardInfo url="jlyon.org"></CardInfo>
    <CardInfo url="s.jlyon.org"></CardInfo>
    <CardInfo url="burndown.io"></CardInfo>
    <CardInfo url="snowsave.me"></CardInfo>




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
