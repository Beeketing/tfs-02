<template>
  <product></product>
  <cart-item></cart-item>
</template>

<script>
import Product from './components/Product.vue';
import CartItem from './components/CartItem.vue';

export default {
  name: "App",
  components: {
    Product,
    CartItem
  }
};
</script>
