<template>
  <div>
    <ScCard
      v-for="scCard in scList"
      :key="scCard.sc.id"
      :name="scCard.sc.name"
      :id="scCard.sc.id"
      :description="scCard.sc.description"
      :phone="scCard.sc.phone"
      :email="scCard.sc.email"
      :site="scCard.sc.site"
      :location="scCard.sc.location"
      :review-info="scCard.review_info"
      @click="openSc(scCard.sc.id)"
    />
  </div>
</template>

<script lang="ts" setup>
import {useScPaginatorStore} from "~/stores/scPaginatorStore";
import {storeToRefs} from "pinia";

const router = useRouter()

const openSc = (id) => {
  //TODO
  router.push(`/spb/sc/${id}-service-center`)
}

const {scList} = storeToRefs(useScPaginatorStore())

</script>

<style scoped>

</style>