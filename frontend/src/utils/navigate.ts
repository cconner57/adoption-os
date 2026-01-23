import { type Router,useRouter } from 'vue-router'

export const goToAdopt = (router: Router, id?: string) => {
  return id ? router.push({ name: 'adopt-pet', params: { id } }) : router.push('/adopt')
}

export const goToDonate = (router: Router) => {
  return router.push('/donate')
}

export const createNavigator = (router: Router) => {
  return {
    goToAdopt: (id?: string) => goToAdopt(router, id),
    goToDonate: () => goToDonate(router),
  }
}

export function useNavigate() {
  const router = useRouter()
  return createNavigator(router)
}
