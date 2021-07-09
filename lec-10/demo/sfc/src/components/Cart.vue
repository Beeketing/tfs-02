<template>
  <div class="row">
    <img
      :src="item.img"
      alt="item.title"
    >
    <h4>{{ item.title }}</h4>
    <p>$ {{ item.price }}</p>
    <div
      class="qty-minus"
      @click="minusQty"
    >
      -
    </div>
    <div class="qty">
      {{ item.qty }}
    </div>
    <div
      class="qty-plus"
      @click="$emit('plusQty')"
    >
      +
    </div>
    <div
      class="del"
      @click="$emit('remove')"
    >
      Remove
    </div>
    <div class="total-price">
      {{ item.price * item.qty }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'CartItem',
  props: {
    item: {
      type: Object,
      default: () => {},
    },
  },
  methods: {
    minusQty() {
      if (this.item.qty > 1) {
        this.$emit('minusQty');
      }
    },
  },
};
</script>
