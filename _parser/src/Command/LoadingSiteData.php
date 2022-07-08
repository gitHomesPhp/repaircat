<?php

declare(strict_types=1);

namespace App\Command;

use App\Service\GisSpider\Review\ReviewSpider;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputArgument;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'load:site:data')]
class LoadingSiteData extends Command
{
    public function __construct(
        private ReviewSpider $reviewSpider,
    ) {
        parent::__construct();
    }

    protected function configure()
    {
        $this->addArgument('project_sc_id', InputArgument::REQUIRED);
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $this->reviewSpider->handle((int)$input->getArgument('project_sc_id'));

        return Command::SUCCESS;
    }
}