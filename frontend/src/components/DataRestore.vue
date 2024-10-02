<template>
  <div>
    <el-autocomplete
      v-model="state2"
      :fetch-suggestions="querySearch"
      :trigger-on-focus="false"
      clearable
      class="inline-input w-50"
      placeholder="Please Enter Data"
      @select="handleSelect"
    />
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";

interface RestaurantItem {
  value: string;
}

const state1 = ref("");
const state2 = ref("");

const restaurants = ref<RestaurantItem[]>([]);
const querySearch = (queryString: string, cb: any) => {
  const results = queryString
    ? restaurants.value.filter(createFilter(queryString))
    : restaurants.value;
  // call callback function to return suggestions
  cb(results);
};
const createFilter = (queryString: string) => {
  return (restaurant: RestaurantItem) => {
    return (
      restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0
    );
  };
};
const loadAll = () => {
  return [{ value: "Data-1" }, { value: "Data-2" }, { value: "Data-3" }];
};

const handleSelect = (item: RestaurantItem) => {
  console.log(item);
};

onMounted(() => {
  restaurants.value = loadAll();
});
</script>
