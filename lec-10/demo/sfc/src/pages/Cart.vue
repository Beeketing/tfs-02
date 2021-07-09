<template>
  <div class="cart">
    <div class="head">
      <h3>Shopping Cart</h3>
      <div class="price">
        Price
      </div>
      <div class="quantity">
        Quantity
      </div>
      <div class="total">
        Total
      </div>
    </div>
    <template v-if="$store.state.cart.items.length > 0">
      <cart
        v-for="(item, index) in $store.state.cart.items"
        :key="index"
        :item="item"
      />
      <h5>SubTotal Amount: $ {{ $store.getters['cart/totalAmount'] }}</h5>
    </template>
  </div>
</template>
<script>
import Cart from '../components/Cart.vue';

export default {
  name: 'PageCart',
  components: {
    Cart,
  }
};
</script>
