<template>
  <div class="main">
    <ul>
      <li>
        <router-link :to="{ name: 'home' }">
          Home
        </router-link>
      </li>
      <li>
        <router-link :to="{ name: 'product', params: { handle: 't-shirt-2021' } }">
          T-shirt
        </router-link>
      </li>
      <li>
        <router-link :to="{ name: 'product', params: { handle: 'adidas' }, query: { variant_id: 1 } }">
          Shoes
        </router-link>
      </li>
    </ul>
    <router-view />
  </div>
</template>
<script>
export default {
  name: 'App',
};
</script>
