<template>
  <div class="paginator" :class="type + '-paginator'">
    <div class="paginator__wrapper">
      <span>
        <a v-if="previous" class="pointer" @click.prevent="previousPage">Предыдущая</a>
      </span>
      <span class="paginator__page">{{ page }}</span>
      <span>
        <a v-if="next" class="pointer" @click.prevent="nextPage">Следующая</a>
      </span>
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
a {
  text-decoration: none;
  color: #0070c0;
}
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