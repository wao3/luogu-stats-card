<template>
  <div id="app">
    <Homepage/>
  </div>
</template>

<script>
import Homepage from './Homepage.vue'

export default {
  name: 'App',
  components: {
    Homepage
  }
}
</script>

<style>
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button{
	-webkit-appearance: none !important;
	margin: 0;
}
#app {
  max-width: 1000px;
  margin: 50px auto;
  padding: 20px
}
</style>
