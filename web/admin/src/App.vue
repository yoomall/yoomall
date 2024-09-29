<template>
  <div>
    <RouterView></RouterView>
    <div :class="defaultUnoClass"></div>
  </div>
</template>
<script>
import { RouterView } from 'vue-router';
export default {
  components: { RouterView },
  props: {},
  data() {
    return {
      defaultUnoClass: [
        "bg-primary",
        "text-primary text-red",
        "text-sm text-xs text-lg text-xl text-2xl text-3xl",
        "flex flex-row flex-col items-center items-start items-end justify-around justify-end justify-start justify-between gap-1 gap-2 gap-4 grid grid-cols-1 grid-cols-2 grid-cols-3 grid-cols-4 grid-cols-5 grid-cols-6",
        "box-border",
        "text-center text-left text-right",
        "mr-5 ml-5",
        "w-80px h-80px w-60px h-60px w-50px h-50px",
        "font-bold font-normal font-light",
        "fixed left-1000% top-1000%",
        "line-clamp-1 line-clamp-2 line-clamp-3 line-clamp-4 line-clamp-5 line-clamp-6",
        "font-normal font-bold font-italic line-through underline underline-offset-1 underline-offset-2 underline-offset-4 underline-offset-8",
      ],
    };
  },
  setup() {

  },
  watch: {},
  computed: {},
  methods: {},
  created() { },
  mounted() { }
};
</script>
<style lang="scss" scoped></style>