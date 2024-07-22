import React from 'react';
import DefaultRouter from './routers/DefaultRouter';
import HomePage from './pages/HomePage';
import NotFoundPage from './pages/NotFoundPage';
import CreateAccountPage from './pages/CreateAccount';
import UserHomePage from './pages/UserHome';
import ChallengeHomePage from './pages/ChallengeHome';
import SubmissionHomePage from './pages/SubmissionHome';
import UserChallengePage from './pages/UserChallenges';
import EasyChallengePage from './pages/EasyChallenges';
import MediumChallengePage from './pages/MediumChallenges';
import HardChallengePage from './pages/HardChallenges';
import CreateChallengePage from './pages/CreateChallenge';
import ViewChallengePage from './pages/ViewChallenge';
import EditChallengePage from './pages/EditChallenge';
import LeaderboardPage from './pages/Leaderboard';

const routes = [
  { path: '/', component: HomePage },
  { path: '*', component: NotFoundPage },
  { path: '/create-account', component: CreateAccountPage },
  { path: '/user-home', component: UserHomePage },
  { path: '/challenge-home', component: ChallengeHomePage },
  { path: '/submission-home', component: SubmissionHomePage },
  { path: '/user-challenges', component: UserChallengePage },
  { path: '/easy-challenges', component: EasyChallengePage },
  { path: '/medium-challenges', component: MediumChallengePage },
  { path: '/hard-challenges', component: HardChallengePage },
  { path: '/create-challenge', component: CreateChallengePage },
  { path: '/view-challenge', component: ViewChallengePage },
  { path: '/edit-challenge', component: EditChallengePage },
  { path: '/leaderboard', component: LeaderboardPage },
];

function App() {
  return (
    <DefaultRouter routes={routes} />
  );
}

export default App;
