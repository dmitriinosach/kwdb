import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
  vus: 2,
  duration: '30s',
  cloud: {
    projectID: 3753077,
    // Test runs with the same name groups test runs together
    name: 'YOUR TEST NAME'
  },
};

export default function() {

  let val = "Наше дело не так однозначно, как может показаться: сплочённость команды профессионалов прекрасно подходит для реализации как самодостаточных, так и внешне зависимых концептуальных решений. Равным образом, начало повседневной работы по формированию позиции представляет собой интересный эксперимент проверки прогресса профессионального сообщества. В своём стремлении повысить качество жизни, они забывают, что базовый вектор развития играет определяющее значение для глубокомысленных рассуждений."
  let vals = "12313t dcvs ga"
  let res = http.get('http://192.168.1.96:713/?key='+getRandomInt(200000)+'&value="'+encodeURIComponent(val)+'"');

  check(res, { "status is 200": (res) => res.status === 200 });

  sleep(1);


}

function getRandomInt(max) {
  return Math.floor(Math.random() * max);
}