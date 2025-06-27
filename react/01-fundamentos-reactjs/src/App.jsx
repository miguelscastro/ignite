import { Header } from './componentes/Header';
import { Sidebar } from './componentes/Sidebar.jsx';
import { Post } from './componentes/Post.jsx';

import styles from './App.module.css';
import './global.css';


/*
    author: { avatar_url: "", name: "", role: "" }
    publishedAt: Date
    content: String
*/

const posts = [
  {
    id: 1,
    author: {
      avatarUrl: "https://github.com/miguelscastro.png",
      name: "Miguel Castro",
      role: "Web Developer"
    },
    content: [
      { type: "paragraph", content: "Fala galeraa 👋" },
      { type: "paragraph", content: "Acabei de subir mais um projeto no meu portifa. É um projeto que fiz no NLW Return, evento da Rocketseat. O nome do projeto é DoctorCare 🚀" },
      { type: "link", content: "👉 miguelscastro/doctorcare" },
    ],
    publishedAt: new Date('2024-01-30 18:45:00'),
  },
  {
    id: 2,
    author: {
      avatarUrl: "https://github.com/castroogbs.png",
      name: "Gabriel Castro",
      role: "Business Architect"
    },
    content: [
      { type: "paragraph", content: "Fala pessoal 👋" },
      { type: "paragraph", content: "Finalmente finalizei meu novo site/portfólio. Foi um baita desafio criar todo o design e codar na unha, mas consegui 💪🏻" },
      { type: "link", content: "👉 castroogbs.design" },
    ],
    publishedAt: new Date('2024-01-31 20:00:00'),
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
                author={post.author}
                content={post.content}
                publishedAt={post.publishedAt}
              />
            )
          })}
        </main>
      </div>
    </div>
  );
}
