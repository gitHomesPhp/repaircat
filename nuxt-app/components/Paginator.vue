<template>
  <div class="paginator" :class="type + '-paginator'">
    <div class="paginator__wrapper">
      <span><span v-if="previous" class="pointer" @click="previousPage">Предыдущая</span></span>
      <span class="paginator__page">{{ page }}</span>
      <span><span v-if="next" class="pointer" @click="nextPage">Следующая</span></span>
    </div>
  </div>
</template>

<script setup lang="ts">
  const props = defineProps({
    type: { type: String },
    page: { type: Number, default: 1},
    next: {},
    previous: {},
  })
  const emit = defineEmits([
      'next-page',
      'previous-page',
  ])

  const nextPage = () => {
      emit('next-page')
  }
  const previousPage = () => {
    if (props.page > 1) {
      emit('previous-page')
    }
  }

</script>

<style scoped lang="scss">
.paginator {
  display: flex;
  &__wrapper {
    width: 100%;
    display: flex;
    justify-content: space-between;
  }
  &__page {
    font-weight: bolder;
  }
}
.top-paginator {
  margin-bottom: 1rem;
  padding-bottom: 1rem;
  border-bottom: #cccccc 1px solid;
}
.bottom-paginator {
  margin-top: 1rem;
  padding-top: 1rem;
  border-top: #cccccc 1px solid;
}
</style>