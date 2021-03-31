import Projects from './content/Projects.svelte';
import ProjectInfo from './content/ProjectInfo.svelte';

export default {
  '/': Projects,
  '/projects/:id': ProjectInfo
}
