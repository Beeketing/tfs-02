export default {
  state: {
    products: []
  },
  mutations: {
    setProducts(state, payload) {
      state.products = payload;
    }
  },
  actions: {
    fetchProducts(context) {
      const products = [
        {
          img: 'https://chenyiya.com/codepen/product-1.jpg',
          title: 'Beer Bottle',
          price: 25,
          id: 1,
        },
        {
          img: 'https://chenyiya.com/codepen/product-2.jpg',
          title: 'Eco Bag',
          price: 73,
          id: 2,
        },
        {
          img: 'https://chenyiya.com/codepen/product-3.jpg',
          title: 'Paper Bag',
          price: 35,
          id: 3,
        },
      ]

      context.commit('setProducts', products);
    }
  }
}
