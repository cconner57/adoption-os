
export function useScrollReveal(className = 'reveal', threshold = 0.1) {
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add('active')
          observer.unobserve(entry.target) 
        }
      })
    },
    {
      threshold,
      rootMargin: '0px 0px -50px 0px', 
    },
  )

  const vScrollReveal = {
    mounted: (el: HTMLElement) => {
      el.classList.add(className)
      observer.observe(el)
    },
    unmounted: (el: HTMLElement) => {
      observer.unobserve(el)
    },
  }

  return {
    vScrollReveal,
  }
}
