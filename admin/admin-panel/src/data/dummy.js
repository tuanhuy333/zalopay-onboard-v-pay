import React from 'react';
import { AiOutlineShoppingCart } from 'react-icons/ai';
import { FiShoppingBag} from 'react-icons/fi';
import { BsCurrencyDollar} from 'react-icons/bs';


export const links = [
  {
    title: 'Dashboard',
    links: [
      {
        name: 'ecommerce',
        icon: <FiShoppingBag />,
      },
    ],
  },

  {
    title: 'Pages',
    links: [
      {
        name: 'orders',
        icon: <AiOutlineShoppingCart />,
      },
      {
        name: 'Item 1',
        icon: <AiOutlineShoppingCart />,
      },
      {
        name: 'Item 2',
        icon: <AiOutlineShoppingCart />,
      },
      {
        name: 'Item 3',
        icon: <AiOutlineShoppingCart />,
      },
    ],
  },
 
];

export const themeColors = [
  {
    name: 'blue-theme',
    color: '#1A97F5',
  },
  {
    name: 'green-theme',
    color: '#03C9D7',
  },
  {
    name: 'purple-theme',
    color: '#7352FF',
  },
  {
    name: 'red-theme',
    color: '#FF5C8E',
  },
  {
    name: 'indigo-theme',
    color: '#1E4DB7',
  },
  {
    color: '#FB9678',
    name: 'orange-theme',
  },
];

export const userProfileData = [
  {
    icon: <BsCurrencyDollar />,
    title: 'My Profile',
    desc: 'Account Settings',
    iconColor: '#03C9D7',
    iconBg: '#E5FAFB',
  },

];
