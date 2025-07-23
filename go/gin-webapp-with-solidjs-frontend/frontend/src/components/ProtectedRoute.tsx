import { ParentComponent, onMount } from 'solid-js';
import { useNavigate } from '@solidjs/router';
import { useAuth } from '../contexts/AuthContext';

export const ProtectedRoute: ParentComponent = (props) => {
  const navigate = useNavigate();
  const { isAuthenticated } = useAuth();

  onMount(() => {
    // Check authentication on mount
    if (!isAuthenticated()) {
      console.log("Not authenticated, redirecting to login");
      navigate('/login');
    } else {
      console.log("User is authenticated, showing protected content");
    }
  });

  // If not authenticated, don't render anything (redirect will happen)
  if (!isAuthenticated()) {
    return null;
  }

  // Render protected content
  return <>{props.children}</>;
};