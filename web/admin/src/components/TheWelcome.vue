<script setup lang="js">
import WelcomeItem from './WelcomeItem.vue'
import DocumentationIcon from './icons/IconDocumentation.vue'
import ToolingIcon from './icons/IconTooling.vue'
import EcosystemIcon from './icons/IconEcosystem.vue'
import CommunityIcon from './icons/IconCommunity.vue'
import SupportIcon from './icons/IconSupport.vue'

import { useProfileStore } from '@/pinia/profile'

const profileStore = useProfileStore()
</script>

<template>
  <div class="">
    <div class="grid xl:grid-cols-4 md:grid-cols-2 sm:grid-cols-1 gap-2">
        <WelcomeItem>
          <template #icon>
            <DocumentationIcon />
          </template>
          <template #heading>
            <span t="t_welcome.documentation">
              {{$t('welcome.documentation',{user:profileStore.profile.username})}}
            </span>
          </template>
          {{$t('welcome.documentationText')}}
        </WelcomeItem>

        <WelcomeItem>
          <template #icon>
            <EcosystemIcon />
          </template>
          <template #heading>
            {{$t('welcome.ecosystem')}}
          </template>
          {{$t('welcome.ecosystemText')}}
        </WelcomeItem>

        <WelcomeItem>
          <template #icon>
            <CommunityIcon />
          </template>
          <template #heading>
            {{$t('welcome.community')}}
          </template>
          {{$t('welcome.communityText')}}
        </WelcomeItem>

        <WelcomeItem>
          <template #icon>
            <SupportIcon />
          </template>
          <template #heading>
            {{$t('welcome.support')}}
          </template>
          {{$t('welcome.supportText')}}
        </WelcomeItem>

    </div>
  </div>
</template>
