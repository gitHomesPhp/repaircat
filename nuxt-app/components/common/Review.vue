<template>
  <div class="review">
    <div class="review__visitor">
      <span class="review__visitor-name">{{ visitor.name }}</span>
      <img src="/img/2gis_icon.svg" width="50" alt="">
      <CommonRating :rating="rating" :count="1" stars-only />
    </div>
    <div class="review__info">
      <div>{{ text }}</div>
    </div>

  </div>
</template>

<script lang="ts" setup>
  const props = defineProps({
    text: { required: true, type: String },
    visitor: { required: true, type: Object },
    rating: { required: true, type: Number },
  })
</script>

<style scoped lang="scss">
.review {
  display: flex;
  margin-bottom: 1rem;
  position: relative;
  &:after {
    position: absolute;
    content: "";
    height: 1px;
    bottom: -.5rem;
    left: -1rem;
    width: calc(100% + 2rem);
    background: #d2d4d9;
  }
  &__visitor {
    padding-right: 1rem;
    display: flex;
    flex-direction: column;
    width: 135px;
    img {
      margin-bottom: .6rem;
    }
  }
  &__info {
    font-family: Roboto,Arial,sans-serif;
    flex-grow: 1;
    display: flex;
    justify-content: flex-start;
  }
  &__visitor-name {
    font-weight: bolder;
    margin-bottom: .4rem;

  }
}
</style>