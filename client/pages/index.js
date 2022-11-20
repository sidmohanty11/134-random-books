import Head from "next/head";
import { useState, useEffect } from "react";

let msg;

export default function Home(props) {
  const [voices, setVoices] = useState([]);

  useEffect(() => {
    msg = new SpeechSynthesisUtterance();
    function getVoices() {
      const voices = speechSynthesis.getVoices();
      setVoices(voices);
    }
    setInterval(() => {
      if (voices.length) {
        clearInterval()
      } else {
        getVoices();
      }
    }, 100)
  }, []);

  const playBook = () => {
    msg.text = props.data.plainText;
    speechSynthesis.speak(msg);
  };

  const setVoice = (e) => {
    msg.voice = voices.find((voice) => voice.name === e.target.value);
  };

  return (
    <div>
      <Head>
        <title>Random 134 Books!</title>
        <meta name="description" content="Get summary of popular 134 books!" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className="container">
        <h1>{props.data.title.toUpperCase()}</h1>
        <button onClick={playBook}>Read it to me!</button>
        <select onChange={setVoice}>
          {voices.map((voice) => (
            <option key={voice.name} value={voice.name}>
              {voice.name} - {voice.lang}
            </option>
          ))}
        </select>
        <div dangerouslySetInnerHTML={{ __html: props.data.book }} />
      </div>
    </div>
  );
}

export const getServerSideProps = async (ctx) => {
  const req = await fetch("http://localhost:8081/handle");
  const data = await req.json();
  return {
    props: data,
  };
};
