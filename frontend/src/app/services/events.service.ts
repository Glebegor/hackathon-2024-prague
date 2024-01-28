import { Injectable } from '@angular/core';
import { group } from '@angular/animations';

export interface ISubscribtion {
  id: number;
  name: string;
  idOfPerson: string;
  price: number;
  description: string;
  image: string;
  tags:string;
  link:string;
}
const subscribtions: ISubscribtion[] = [
  {
    id: 1,
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    image: "https://fakestoreapi.com/img/71-3HjGNDUL._AC_SY879._SX._UX._SY._UY_.jpg",
    tags:"dfsgfaaaaaaaaaaaaaaaaaaa,sgfgfs,sfdgsfdg,sdgfsgf,gssdfg,sfggs,sghdg,gfcn",
    link:"http://google.com"
  },
  {
    id: 2, 
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    image: "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
    tags:"bom",
    link:"http://google.com"
  },
  {
    id: 3,
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    image: "https://fakestoreapi.com/img/71-3HjGNDUL._AC_SY879._SX._UX._SY._UY_.jpg",
    tags:"#bom",
    link:"http://google.com"
  },
  {
    id: 4, 
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    image: "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
    tags:"#bom",
    link:"http://google.com"
  },
  {
    id: 5,
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    image: "https://fakestoreapi.com/img/71-3HjGNDUL._AC_SY879._SX._UX._SY._UY_.jpg",
    tags:"#bom",
    link:"http://google.com"
  },
  {
    id: 6, 
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    image: "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
    tags:"#bom",
    link:"http://google.com"
  },
  {
    id: 7,
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    image: "https://fakestoreapi.com/img/71-3HjGNDUL._AC_SY879._SX._UX._SY._UY_.jpg",
    tags:"#bom",
    link:"http://google.com"
  },
  {
    id: 8, 
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    image: "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
    tags:"#bom",
    link:"http://google.com"
  },
];

``
@Injectable({
  providedIn: 'root'
})
export class EventsService {
  readonly subscribtions: ISubscribtion [] = subscribtions ;

  getById(id: number) {  
    return this.subscribtions.find(subscribtion => subscribtion.id === id);
  }

  getSubscribtions(){
    return this.subscribtions;
  }

  // get byGroup() {
  //   return this.events.reduce((group: { [key: string]: IEvent[] }, event) => {
  //     if (!group[event.category]) {
  //       group[event.category] = [];
  //     }
  //     group[event.category].push(event);
  //     return group;
  //   }, {});
  // }


}
