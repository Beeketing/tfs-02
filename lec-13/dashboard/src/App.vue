<template>
  <div id="app">
    <select name="cars" id="cars" v-model="gateway">
      <option value="stripe">stripe</option>
      <option value="paypal">paypal</option>
    </select>

    <component :is="gateway" ref="gateway"/>
    <div @click="pay">Pay now</div>
  </div>
</template>

<script>
import Stripe from "@/components/Stripe";
import Paypal from "@/components/Paypal";

export default {
  name: 'App',
  components: {
    Paypal,
    Stripe,
  },
  data() {
    return {
      gateway: 'stripe'
    }
  },
  methods: {
    pay() {
      this.$refs.gateway.createPaymentMethod().then(console.log)
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
