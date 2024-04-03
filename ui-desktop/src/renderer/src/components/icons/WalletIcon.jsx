import React from 'react';
import styled from 'styled-components';

import BaseIcon from './BaseIcon';

export const WalletIcon = ({ size, fill = 'black' }, props) => (
  <BaseIcon size={size} viewBox="0 0 40 40" {...props}>
    <path
      fillRule="evenodd"
      clipRule="evenodd"
      d="M23.351 30C23.2019 30.0006 23.0531 29.9838 22.9078 29.95V29.95L12.1848 27.6624C11.2048 27.4555 10.4985 26.5979 10.4834 25.5964V15.5096C10.4834 14.3252 11.4436 13.365 12.628 13.365H28.3551C29.5395 13.365 30.4997 14.3252 30.4997 15.5096V25.5178C30.4997 26.7022 29.5396 27.6624 28.3551 27.6624H25.4956V27.8554C25.4956 29.0398 24.5355 30 23.351 30V30ZM23.2009 28.5703C23.4135 28.614 23.6344 28.5587 23.8014 28.4201C23.9675 28.2854 24.0645 28.0835 24.0659 27.8697V17.9044C24.0735 17.5615 23.8365 17.2615 23.5012 17.1896L12.7781 14.9306C12.5666 14.8849 12.3458 14.9375 12.1776 15.0735C12.0037 15.2145 11.9058 15.4288 11.9131 15.6526V25.5964C11.9055 25.9393 12.1425 26.2393 12.4779 26.3113L23.2009 28.5703ZM25.4956 26.2326H28.3551C28.7499 26.2326 29.07 25.9126 29.07 25.5178V21.0355H25.4956L25.4956 26.2326ZM18.976 14.7947L23.7942 15.8099C24.7846 16.019 25.4939 16.8922 25.4956 17.9044V19.6058H29.07V15.5096C29.07 15.1148 28.7499 14.7947 28.3551 14.7947H18.976ZM20.849 24.6099C20.2568 24.6099 19.7767 24.1298 19.7767 23.5376C19.7767 22.9454 20.2568 22.4653 20.849 22.4653C21.4412 22.4653 21.9213 22.9454 21.9213 23.5376C21.9213 24.1298 21.4412 24.6099 20.849 24.6099ZM20.849 23.1802C20.6516 23.1802 20.4916 23.3402 20.4916 23.5376C20.4916 23.735 20.6516 23.895 20.849 23.895C21.0464 23.895 21.2064 23.735 21.2064 23.5376C21.2064 23.3402 21.0464 23.1802 20.849 23.1802Z"
      fill={fill}
    />
  </BaseIcon>
);
