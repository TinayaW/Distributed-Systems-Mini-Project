import React, { useState } from 'react';

type Route = {
  path: string;
  component: React.ComponentType<any>;
};

type RouterProps = {
  routes: Route[];
};

const DefaultRouter: React.FC<RouterProps> = ({ routes }) => {
  const [currentPath, setCurrentPath] = useState(window.location.pathname);

  window.onpopstate = () => {
    setCurrentPath(window.location.pathname);
  };

  const Component = routes.find(route => route.path === currentPath)?.component;

  return Component ? <Component /> : null;
};

export default DefaultRouter;
