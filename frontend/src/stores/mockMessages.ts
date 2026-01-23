import { ref } from 'vue'

export interface IChatUser {
  id: string
  name: string
  avatar: string
  role: 'admin' | 'volunteer' | 'adopter'
}

export interface IThread {
  id: string
  type: 'channel' | 'inquiry'
  contactMethod?: 'sms' | 'email'
  name: string
  description: string
  icon: string
  isPrivate?: boolean
  members?: string[] 
}

export interface IMessage {
  id: string
  threadId: string
  userId: string
  text: string
  timestamp: string
  type: 'text' | 'system'
}

export const currentUser = ref<IChatUser>({
  id: 'u-admin',
  name: 'Admin User',
  avatar: '👩‍💼',
  role: 'admin',
})

export const mockUsers = ref<IChatUser[]>([
  currentUser.value,
  { id: 'u-1', name: 'Sarah Jenkins', avatar: '👩‍🌾', role: 'volunteer' },
  { id: 'u-2', name: 'Mike Ross', avatar: '👨‍⚖️', role: 'volunteer' },
  { id: 'u-3', name: 'Jessica Pearson', avatar: '👩‍⚖️', role: 'volunteer' },
  { id: 'u-4', name: 'John Doe (Adopter)', avatar: '🧔', role: 'adopter' },
  { id: 'u-5', name: 'Emily Smith (Adopter)', avatar: '👩', role: 'adopter' },
])

export const mockThreads = ref<IThread[]>([
  
  {
    id: 'th-1',
    type: 'channel',
    name: 'general-questions',
    description: 'Help and questions for everyone',
    icon: '#️⃣',
  },
  {
    id: 'th-2',
    type: 'channel',
    name: 'teen-volunteers',
    description: 'Coordination for teen program',
    icon: '🎒',
    isPrivate: true,
  },
  {
    id: 'th-3',
    type: 'channel',
    name: 'cat-photos',
    description: 'Share your best cat pics and memes!',
    icon: '🐱',
  },
  {
    id: 'th-4',
    type: 'channel',
    name: 'announcements',
    description: 'Official shelter updates',
    icon: '📢',
  },
  
  {
    id: 'th-inq-1',
    type: 'inquiry',
    contactMethod: 'sms',
    name: 'John Doe',
    description: 'Inquiry about "Rex"',
    icon: '📱',
  },
  {
    id: 'th-inq-2',
    type: 'inquiry',
    contactMethod: 'email',
    name: 'Emily Smith',
    description: 'Application follow-up',
    icon: '📧',
  },
])

export const mockMessages = ref<IMessage[]>([
  {
    id: 'm-1',
    threadId: 'th-1',
    userId: 'u-1',
    text: 'Hey everyone! Where do we store the extra blankets?',
    timestamp: '10:30 AM',
    type: 'text',
  },
  {
    id: 'm-2',
    threadId: 'th-1',
    userId: 'u-admin',
    text: 'They are in the storage closet next to the break room.',
    timestamp: '10:35 AM',
    type: 'text',
  },
  {
    id: 'm-3',
    threadId: 'th-3',
    userId: 'u-2',
    text: 'Look at this little guy I found today! 🥺',
    timestamp: '11:00 AM',
    type: 'text',
  },
  {
    id: 'm-4',
    threadId: 'th-3',
    userId: 'u-3',
    text: 'Omg so cute!!',
    timestamp: '11:05 AM',
    type: 'text',
  },
  
  {
    id: 'm-5',
    threadId: 'th-inq-1',
    userId: 'u-4', 
    text: 'Hi, I saw Rex on the website. Is he good with cats?',
    timestamp: '09:00 AM',
    type: 'text',
  },
  {
    id: 'm-6',
    threadId: 'th-inq-1',
    userId: 'u-admin',
    text: 'Hello John! Yes, Rex has lived with cats before.',
    timestamp: '09:15 AM',
    type: 'text',
  },
])
