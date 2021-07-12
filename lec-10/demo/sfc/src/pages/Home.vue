<template>
  <div>
    <section class="section-wrap promo-banners pb-30">
      <div class="container">
        <div class="row">
          <PromoBanner
            title="her"
            description="Best Selling Deals"
            :image="bannerOne"
          />
        </div>
      </div>
    </section>
  </div>
</template>
<script>
import PromoBanner  from '../components/PromoBanner.vue';
import bannerOne from '../../assets/img/shop/collection_1.jpg'

export default {
  name: 'Home',
  components: {
    PromoBanner,
  },
  data() {
    return {
      bannerOne
    }
  }
};
</script>
