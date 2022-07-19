<template>
  <div class="review">
    <div class="review__visitor">
      <span class="review__visitor-name">{{ visitor.name }}</span>
      <img src="/img/2gis_icon.svg" width="50" alt="">
    </div>
    <div class="review__info">
      {{ text }}
    </div>
  </div>
</template>

<script lang="ts" setup>
  const props = defineProps({
    text: { required: true, type: String },
    visitor: { required: true, type: Object }
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
    align-items: center;
    width: 20%;
  }
  &__info {
    width: 80%;
  }
  &__visitor-name {
    font-weight: bolder;
    margin-bottom: .4rem;

  }
}
</style>