export default {
  namespaced: true,
  state: {
    items: [
      {
        img: 'https://chenyiya.com/codepen/product-1.jpg',
        title: 'Beer Bottle',
        price: 25,
        id: 1,
        qty: 3,
      },
    ]
  },
  getters: {
    totalAmount(state) {
      return state.items.reduce((currentValue, item) => item.price * item.qty + currentValue, 0);
    }
  }
}
