<template>
  <div class="rating">
    <div class="rating__star">
      <div class="rating__fill" :style="'width:' + ratingFill + '%'"></div>
    </div>
    <div class="rating__count">
      <span>{{ props.count}} отзывов</span>
    </div>
  </div>
</template>

<script lang="ts" setup>
  const props = defineProps({
    count: { type: Number, required: true },
    rating: {type: Number, required: true}
  })

  const ratingFill  = 100 * props.rating / 5
</script>

<style scoped lang="scss">
.rating {
  &__star {
    background: url(/img/star_icon.svg);
    width: 135px;
    height: 24px;
  }
  &__fill {
    background: url(/img/star_icon_fill.svg);
    height: 24px;
  }
  &__count {
    margin-top: 3px;
    font-size: .9rem;
    text-decoration: underline;
    color: #6c757d;
  }
}
</style>