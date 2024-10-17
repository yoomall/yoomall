<script setup>
import { RouterLink, RouterView, useRoute, useRouter } from 'vue-router'
import Menu from '../components/layout/Menu.vue'
import { ElAvatar, ElBreadcrumb, ElButton, ElTabs, ElMessageBox } from 'element-plus';
import { computed, onMounted, ref, watch, watchEffect } from 'vue';
import useTranslateStore from '../pinia/translate';
import { request } from '../api/request';
import useProfileStore from '../pinia/profile';
import { useSubMenuStore } from '../pinia/subMenu';
const menus = ref([])
const router = useRouter()
const translateStore = useTranslateStore()
const profileStore = useProfileStore()
const subMenuStore = useSubMenuStore()

const currentPath = computed(() => {
    return router.currentRoute.value.path
})
onMounted(() => {
    request.get('/menus').then(res => {
        menus.value = (res.data.data || []).filter(el => el.hidden !== true).map(el => {
            if (el.children) {
                el.children = el.children.filter(el => el.hidden !== true)
            }
            return el
        })
    })

    profileStore.refreshProfile()

    findCurrentMenuChildren()

})

const findCurrentMenuChildren = () => {


    if (router.currentRoute.value.matched?.[1]?.meta?.children) {
        subMenuStore.setSubMenu(router.currentRoute.value.matched?.[1]?.meta?.children)
    }

    router.afterEach((to, form, err) => {
        console.log("to", to)
        if (to.matched?.[1]?.meta?.children) {
            subMenuStore.setSubMenu(to.matched?.[1]?.meta?.children)
        } else {
            subMenuStore.setSubMenu([])
        }
    })
}

const logout = () => {
    // confirm 
    ElMessageBox.confirm('Are you sure to logout?', 'Logout', {
        confirmButtonText: 'Logout',
        cancelButtonText: 'Cancel',
        type: 'warning',
    }).then(() => {
        request.post('/auth/logout').then(res => {
            location.href = '/'
        })
    }).catch(() => {
        // cancel
    });

}

</script>

<template>
    <div class="min-h-screen">
        <el-container class="min-h-screen bg-gary-100">
            <div class="light:text-light-header-text light:bg-light-header-bg dark:bg-dark dark:text-white fixed w-full px-4 box-border border-solid border-0 border-b-0px light:border-gray-300 dark:border-dark"
                style="left: 0;top:0;z-index: 99;height: var(--header-bar-height)">
                <div class="flex flex-row items-center py-2">
                    <div class="flex-1">
                        <UIButton>
                            <span class="text-lg !mt-0">{{ $t('AdminTitle') }}</span>
                        </UIButton>
                    </div>
                    <div class="flex flex-row items-center justify-center gap-4">
                        <UIButton class="flex-row-btn">
                            <Icon icon="ant-design:sort-ascending-outlined"></Icon>
                            <span>{{ translateStore.getLocaleToDisplay() }}</span>
                        </UIButton>
                        <UIButton class="flex-row-btn">
                            <Icon icon="ant-design:message-outlined"></Icon>
                            <span>{{ $t('header.notifyMessage') }}</span>
                        </UIButton>
                        <UIButton class="flex-row-btn">
                            <Icon icon="ant-design:setting-outlined"></Icon>
                            <span>{{ $t("setting") }}</span>
                        </UIButton>
                        <UIButton class="flex-row-btn" @click="logout">
                            <Icon icon="ant-design:login-outlined"></Icon>
                            <span>{{ profileStore.profile.username }}</span>
                        </UIButton>
                        <UIButton style="width: 32px;height: 32px;">
                            <ElAvatar style="width: 32px;height: 32px;" :src="$img(profileStore.profile.avatar)" />
                        </UIButton>
                    </div>
                </div>
            </div>
            <div class="flex flex-row flex-1 light:bg-gray-100 h-screen"
                style="padding-top: var(--header-bar-height);box-sizing: border-box;">
                <div class="w-64px dark:bg-dark-800 light:bg-dark light:text-white overflow-y-auto hidden-scroll-bar border-0 border-solid border-r-1px light:border-gray-200 dark:border-dark"
                    :style="subMenuStore.hasSubMenu ? `z-index: 9;box-shadow: 6px 0 36px #00000010;` : ''">
                    <!-- <Menu :menus="menus" ></Menu> -->
                    <div>
                        <div @click="() => {
                            router.push(m.path)
                            subMenuStore.setSubMenu(m.children ?? [])
                        }" :class="{
                        'bg-dark-100 dark:bg-dark-600 on': currentPath.startsWith(m.path)
                    }" v-for="m in menus"
                            class="flex flex-col items-center p-2 py-3 text-center cursor-pointer hover:bg-dark-100 dark:hover:bg-dark-600 hover:text-gray-300">
                            <Icon :icon="m.icon" />
                            <span class="text-sm">{{ m.title }}</span>
                        </div>
                    </div>
                    <div class="h-100px"></div>
                </div>
                <div v-if="subMenuStore.hasSubMenu" class="w-180px bg-white dark:bg-dark-800">
                    <Menu :menus="subMenuStore.menus" is-sub-menu></Menu>
                </div>
                <div class="flex-1 flex flex-col dark:bg-dark-700"
                    style="height:var(100vh - var(--header-bar-height));overflow-y: auto;">
                    <main class="flex-1">

                        <div class="">
                            <div class="p-2 py-0 bg-white dark:bg-dark-600 text-gray-400">
                                DEBUG:{{ currentPath }}
                            </div>
                        </div>

                        <RouterView :key="$route.meta.key"></RouterView>
                    </main>
                </div>
            </div>
            <!-- <el-footer class="bg-dark-600">footer</el-footer> -->
        </el-container>
    </div>
</template>

<style scoped></style>
