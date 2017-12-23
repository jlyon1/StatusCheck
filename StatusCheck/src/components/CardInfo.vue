<template>
  <div class="tile is-child is-primary">
  <div class="card">
  <header class="card-header">
  <p class="card-header-title">
  {{url}}
  </p>
  </header>
  <div class="card-content">

  <div class="notification" v-bind:class="{'is-primary': good, 'is-danger': !good}">
  <div class="level">
  <p v-if="good">
    Live
  </p>
  <p v-if="!good">
    Not Live
  </p>
  </div>
  </div>
  <div class="content">
  <div class="">
  <button @click=go class="button is-fullwidth">Visit</button>
  </div>
  </div>
  </div>

  </div>
  </div>
</template>

<script>
export default {
  name: 'CardInfo',
  props: ['url'],
  data () {
    return {
      good: true,
    }


  },
  methods:
  {
    go: function(){
      window.location.href = "https://" + this.url;
    },
    check: function(){
      let el = this;
      $.get('/Live/' + el.url, function(data){
        el.good = data.Live
      });
    }
  },
  mounted(){
    this.check()
    setInterval(this.check,10000);
  }
}
</script>
