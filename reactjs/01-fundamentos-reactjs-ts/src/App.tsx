import { Header } from './componentes/Header';
import { Sidebar } from './componentes/Sidebar';
import { Post, PostType } from './componentes/Post';

import styles from './App.module.css';
import './global.css';


/*
    author: { avatar_url: "", name: "", role: "" }
    publishedAt: Date
    content: String
*/

const posts: PostType[] = [
  {
    id: 1,
    author: {
      avatarUrl: "/post-pfp-1.png",
      name: "Jane Cooper",
      role: "Dev Front-End"
    },
    content: [
      { type: "paragraph", content: "Fala galeraa 👋" },
      { type: "paragraph", content: "Acabei de subir mais um projeto no meu portifa. É um projeto que fiz no NLW Return, evento da Rocketseat. O nome do projeto é DoctorCare 🚀" },
      { type: "link", content: "👉 jane.design/doctorcare" },
    ],
    publishedAt: new Date('2025-01-30 18:45:00'),
  },
  {
    id: 2,
    author: {
      avatarUrl: "/post-pfp-2.png",
      name: "Devon Lane",
      role: "Dev Front-End"
    },
    content: [
      { type: "paragraph", content: "Fala pessoal 👋" },
      { type: "paragraph", content: "Finalmente finalizei meu novo site/portfólio. Foi um baita desafio criar todo o design e codar na unha, mas consegui 💪🏻" },
      { type: "link", content: "👉 devonlane.design" },
    ],
    publishedAt: new Date('2025-01-03 20:00:00'),
  },

]

export function App() {
  return (
    <div>
      <Header />

      <div className={styles.wrapper}>
        <Sidebar />
        <main>
          {posts.map(post => {
            return (
              <Post
                key={post.id}
                post={post}
              />
            )
          })}
        </main>
      </div>
    </div>
  );
}
